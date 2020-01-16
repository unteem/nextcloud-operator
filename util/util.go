package util

import (
	appsv1beta1 "git.indie.host/nextcloud-operator/api/v1beta1"
	version "github.com/hashicorp/go-version"
)

//. TOFIX get install status and desired status, if status =! nil ??
func GetAppPhase(status appsv1beta1.NextcloudStatus, desiredVersion string) (appsv1beta1.Phase, error) {
	// get nextcloud installed version in status
	if len(status.Version) == 0 {
		return appsv1beta1.PhaseInstalling, nil
	}
	v1, err := version.NewVersion(status.Version)
	// get nextcloud desired version in spec
	v2, err := version.NewVersion(desiredVersion)
	if err != nil {
		return appsv1beta1.PhaseFailed, err
	}
	// compare current version from status with desired from spec
	c := v1.Compare(v2)
	switch {
	// current version is greater than desired, no downgrade
	case c == 1:
		return appsv1beta1.PhaseFailed, nil
	// current version lower than desired, upgrade job
	case c == -1:
		return appsv1beta1.PhaseUpgrading, nil
	// current version and desired are the same, normal start
	case c == 0:
		return appsv1beta1.PhaseCreating, nil
	}
	return appsv1beta1.PhaseNone, nil
}
