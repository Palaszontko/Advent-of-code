set dotenv-load
set quiet

YEAR := "2024"
new-day DAY: 
    echo "Advent of code!"  
    mkdir -p 'cmd/{{YEAR}}/day_{{DAY}}'
    cp -R template/main.go cmd/{{YEAR}}/day_{{DAY}}/main.go
    curl --cookie "session=$SESSION" https://adventofcode.com/{{YEAR}}/day/{{DAY}}/input -o cmd/{{YEAR}}/day_{{DAY}}/input.txt

run DAY:
    echo "Advent of code - day {{DAY}}"
    go run cmd/{{YEAR}}/day_{{DAY}}/main.go

remove DAY:
    rm -rf cmd/{{YEAR}}/day_{{DAY}}