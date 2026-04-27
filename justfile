set dotenv-load
set quiet

new-day YEAR DAY: 
    echo "Advent of code!"  
    mkdir -p 'cmd/{{YEAR}}/day_{{DAY}}'
    cp -R template/main.go cmd/{{YEAR}}/day_{{DAY}}/main.go
    sed -i '' 's/Advent of Code [0-9]*/Advent of Code {{YEAR}}/g' cmd/{{YEAR}}/day_{{DAY}}/main.go
    curl --cookie "session=$SESSION" https://adventofcode.com/{{YEAR}}/day/{{DAY}}/input -o cmd/{{YEAR}}/day_{{DAY}}/input.txt

run YEAR DAY:
    echo "Advent of code - day {{DAY}}"
    go run cmd/{{YEAR}}/day_{{DAY}}/main.go

input YEAR DAY:
    echo "Getting input for Advent of Code {{YEAR}} - day {{DAY}}"
    curl --cookie "session=$SESSION" https://adventofcode.com/{{YEAR}}/day/{{DAY}}/input -o cmd/{{YEAR}}/day_{{DAY}}/input.txt

remove YEAR DAY:
    rm -rf cmd/{{YEAR}}/day_{{DAY}}

new-day-rs YEAR DAY:
    echo "Advent of code (Rust)!"
    mkdir -p 'rust/{{YEAR}}/day_{{DAY}}/src' 'rust/{{YEAR}}/day_{{DAY}}/benches'
    cp rust/template/src/main.rs rust/{{YEAR}}/day_{{DAY}}/src/main.rs
    cp rust/template/src/lib.rs rust/{{YEAR}}/day_{{DAY}}/src/lib.rs
    cp rust/template/benches/bench.rs rust/{{YEAR}}/day_{{DAY}}/benches/bench.rs
    cp rust/template/Cargo.toml.tmpl rust/{{YEAR}}/day_{{DAY}}/Cargo.toml
    sed -i '' 's/__YEAR__/{{YEAR}}/g; s/__DAY__/{{DAY}}/g' \
        rust/{{YEAR}}/day_{{DAY}}/Cargo.toml \
        rust/{{YEAR}}/day_{{DAY}}/src/main.rs \
        rust/{{YEAR}}/day_{{DAY}}/src/lib.rs \
        rust/{{YEAR}}/day_{{DAY}}/benches/bench.rs
    grep -q '"{{YEAR}}/day_{{DAY}}"' rust/Cargo.toml || (awk -v m='    "{{YEAR}}/day_{{DAY}}",' '/# AOC_MEMBERS_MARKER/ && !done {print m; done=1} {print}' rust/Cargo.toml > rust/Cargo.toml.tmp && mv rust/Cargo.toml.tmp rust/Cargo.toml)
    curl --cookie "session=$SESSION" https://adventofcode.com/{{YEAR}}/day/{{DAY}}/input -o rust/{{YEAR}}/day_{{DAY}}/input.txt

run-rs YEAR DAY:
    echo "Advent of code (Rust) - day {{DAY}}"
    cargo run --release --manifest-path rust/Cargo.toml -p aoc-{{YEAR}}-day-{{DAY}}

bench-rs YEAR DAY:
    echo "Benchmarking Advent of Code {{YEAR}} - day {{DAY}}"
    cargo bench --manifest-path rust/Cargo.toml -p aoc-{{YEAR}}-day-{{DAY}} --bench bench

time-rs YEAR DAY:
    echo "==> compile (release, clean):"
    cargo clean --manifest-path rust/Cargo.toml -p aoc-{{YEAR}}-day-{{DAY}}
    time cargo build --release --manifest-path rust/Cargo.toml -p aoc-{{YEAR}}-day-{{DAY}}
    echo ""
    echo "==> run (release):"
    time cargo run --release --manifest-path rust/Cargo.toml -p aoc-{{YEAR}}-day-{{DAY}}

input-rs YEAR DAY:
    echo "Getting input for Advent of Code {{YEAR}} - day {{DAY}} (Rust)"
    curl --cookie "session=$SESSION" https://adventofcode.com/{{YEAR}}/day/{{DAY}}/input -o rust/{{YEAR}}/day_{{DAY}}/input.txt

remove-rs YEAR DAY:
    rm -rf rust/{{YEAR}}/day_{{DAY}}
    sed -i '' '\|"{{YEAR}}/day_{{DAY}}"|d' rust/Cargo.toml