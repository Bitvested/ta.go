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
// The Go module is tagged in the v0.x series for module-path compatibility,
// but the user-facing version (matching the sister ta.js library) is recorded
// in [Version].
package ta

// Version is the user-facing release version of this library, kept in sync
// with the matching ta.js release. The Go module tag itself stays in the
// v0.x series — see docs/v2-upgrade-plan.md for the rationale.
const Version = "2.0.0"
