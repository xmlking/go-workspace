/*
Package version implements BuildInfo

Usage:

    log.Info().Object("build_info", version.GetBuildInfo()).Send()
    log.Info().Msgf("build_info:%s", version.GetBuildInfo().PrettyString())
    log.Info().Msg(version.GetSoftwareBOM())
*/
package version
