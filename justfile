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

remove YEAR DAY:
    rm -rf cmd/{{YEAR}}/day_{{DAY}}