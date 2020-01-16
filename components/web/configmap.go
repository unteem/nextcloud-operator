/*

Licensed under the GNU AFFERO GENERAL PUBLIC LICENSE, Version 3 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://www.gnu.org/licenses/agpl-3.0.html

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package web

import (
	"bytes"
	"html/template"

	"github.com/presslabs/controller-util/syncer"

	interfaces "git.indie.host/nextcloud-operator/interfaces"
)

const conf = `
	user www-data;
	events {
	worker_connections 768;
	}

	http {
	upstream backend {
		server {{ .Name }}-app:9000;
	}
	include /etc/nginx/mime.types;
	default_type application/octet-stream;

	server {
		listen 80;

		# Add headers to serve security related headers
		add_header X-Content-Type-Options nosniff;
		add_header X-XSS-Protection "1; mode=block";
		add_header X-Robots-Tag none;
		add_header X-Download-Options noopen;
		add_header X-Permitted-Cross-Domain-Policies none;
		add_header Referrer-Policy no-referrer;

		root /var/www/html;

		location = /robots.txt {
		allow all;
		log_not_found off;
		access_log off;
		}

		location = /.well-known/carddav {
		return 301 https://$host/remote.php/dav;
		}
		location = /.well-known/caldav {
		return 301 https://$host/remote.php/dav;
		}

		client_max_body_size 1G;
		fastcgi_buffers 64 4K;

		gzip off; # handled at haproxy level

		location / {
			rewrite ^ /index.php;
		}

		location ~ ^\/(?:build|tests|config|lib|3rdparty|templates|data)\/ {
			deny all;
		}

		location ~ ^\/(?:\.|autotest|occ|issue|indie|db_|console) {
			deny all;
		}

		location ~ ^\/(?:index|remote|public|cron|core\/ajax\/update|status|ocs\/v[12]|updater\/.+|oc[ms]-provider\/.+)\.php(?:$|\/) {
			fastcgi_split_path_info ^(.+\.php)(/.*)$;
			try_files $fastcgi_script_name =404;
			set $path_info $fastcgi_path_info;
			include fastcgi_params;
			fastcgi_param SCRIPT_FILENAME $document_root$fastcgi_script_name;
			fastcgi_param PATH_INFO $path_info;
			fastcgi_param HTTPS on;
			#Avoid sending the security headers twice
			fastcgi_param modHeadersAvailable true;
			fastcgi_param front_controller_active true;
			fastcgi_pass backend;
			fastcgi_intercept_errors on;
			fastcgi_request_buffering off;
		}

		location ~ ^\/(?:updater|oc[ms]-provider)(?:$|\/) {
			try_files $uri/ =404;
			index index.php;
		}

		# Adding the cache control header for js and css files
		# Make sure it is BELOW the PHP block
		location ~ \.(?:css|js|woff2?|svg|gif|map)$ {
			try_files $uri /index.php$request_uri;
			add_header Cache-Control "public, max-age=15778463";
			# Add headers to serve security related headers (It is intended to
			# have those duplicated to the ones above)
			# Before enabling Strict-Transport-Security headers please read into
			# this topic first.
			# add_header Strict-Transport-Security "max-age=15768000;
			#  includeSubDomains; preload;";
			add_header X-Content-Type-Options nosniff;
			add_header X-Frame-Options "SAMEORIGIN";
			add_header X-XSS-Protection "1; mode=block";
			add_header X-Robots-Tag none;
			add_header X-Download-Options noopen;
			add_header X-Permitted-Cross-Domain-Policies none;
			add_header Referrer-Policy no-referrer;
			# Optional: Don't log access to assets
			access_log off;
		}

		location ~ \.(?:png|html|ttf|ico|jpg|jpeg|bcmap)$ {
			try_files $uri /index.php$request_uri;
			# Optional: Don't log access to other assets
			access_log off;
		}
	}
}
`

func (component *Component) NewConfigMapSyncer(r interfaces.Reconcile) syncer.Interface {
	return syncer.NewObjectSyncer("ConfigMap", component.Owner, &component.ConfigMap, r.GetClient(), r.GetScheme(), component.MutateConfigMap)
}

func (component *Component) MutateConfigMap() error {
	labels := component.Labels("web")

	component.ConfigMap.SetLabels(labels)
	data, err := component.GenConfigMapData()
	if err != nil {
		return err
	}
	component.ConfigMap.Data = data

	return nil
}

// GenAppConfigMapData func generates data for the web configmap that contains the nginx.conf
func (c *Component) GenConfigMapData() (map[string]string, error) {
	var cm map[string]string
	tmpl, err := template.New("test").Parse(conf)
	if err != nil {
		return nil, err
	}
	var tpl bytes.Buffer
	err = tmpl.Execute(&tpl, c.Owner)
	if err != nil {
		return nil, err
	}

	cm = map[string]string{
		"nginx.conf": tpl.String(),
	}

	return cm, nil
}
