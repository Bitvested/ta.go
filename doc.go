// Package ta is the documentation root for the ta.go technical-analysis
// library. The actual implementations live in category subpackages under src/:
//
//   - github.com/Bitvested/ta.go/src/moving_averages
//   - github.com/Bitvested/ta.go/src/indicators
//   - github.com/Bitvested/ta.go/src/oscillators
//   - github.com/Bitvested/ta.go/src/bands
//   - github.com/Bitvested/ta.go/src/statistics
//   - github.com/Bitvested/ta.go/src/chart_types
//   - github.com/Bitvested/ta.go/src/experimental
//   - github.com/Bitvested/ta.go/src/random
//   - github.com/Bitvested/ta.go/src/misc
//
// The Go module is tagged in the v0.x series for module-path compatibility;
// the user-facing release version is recorded in [Version].
//
// # Maintained By
//
// ta.go is an open-source project created and actively maintained by
// Nino Kroesen (https://ninokroesen.com).
//
// Built with performance in mind for high-frequency data and algorithmic
// trading systems — see Bitvested (https://bitvested.com) for these
// indicators in a live quantitative trading environment.
package ta

// Version is the user-facing release version of this library. The Go
// module tag itself stays in the v0.x series — see
// docs/v2-upgrade-plan.md for the rationale.
const Version = "2.0.0"
