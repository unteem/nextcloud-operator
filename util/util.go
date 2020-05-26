package util

import (
	appsv1alpha1 "git.indie.host/operators/nextcloud-operator/api/v1alpha1"
	version "github.com/hashicorp/go-version"
)

//. TOFIX get install status and desired status, if status =! nil ??
func GetAppPhase(status appsv1alpha1.NextcloudStatus, desiredVersion string) (appsv1alpha1.Phase, error) {
	// get nextcloud installed version in status
	if len(status.Version) == 0 {
		return appsv1alpha1.PhaseInstalling, nil
	}
	v1, err := version.NewVersion(status.Version)
	// get nextcloud desired version in spec
	v2, err := version.NewVersion(desiredVersion)
	if err != nil {
		return appsv1alpha1.PhaseFailed, err
	}
	// compare current version from status with desired from spec
	c := v1.Compare(v2)
	switch {
	// current version is greater than desired, no downgrade
	case c == 1:
		return appsv1alpha1.PhaseFailed, nil
	// current version lower than desired, upgrade job
	case c == -1:
		return appsv1alpha1.PhaseUpgrading, nil
	// current version and desired are the same, normal start
	case c == 0:
		return appsv1alpha1.PhaseCreating, nil
	}
	return appsv1alpha1.PhaseNone, nil
}
