set dotenv-load
set quiet

new-day-go YEAR DAY:
    echo "Advent of code!"
    mkdir -p 'cmd/{{YEAR}}/day_{{DAY}}'
    cp -R template/main.go cmd/{{YEAR}}/day_{{DAY}}/main.go
    sed -i '' 's/Advent of Code [0-9]*/Advent of Code {{YEAR}}/g' cmd/{{YEAR}}/day_{{DAY}}/main.go
    curl --cookie "session=$SESSION" https://adventofcode.com/{{YEAR}}/day/{{DAY}}/input -o cmd/{{YEAR}}/day_{{DAY}}/input.txt

new-day-java YEAR DAY:
    echo "Advent of code!"
    mkdir -p 'cmd/{{YEAR}}/day_{{DAY}}'
    cp -R template/Main.java cmd/{{YEAR}}/day_{{DAY}}/Main.java
    sed -i '' 's/Advent of Code [0-9]*/Advent of Code {{YEAR}}/g' cmd/{{YEAR}}/day_{{DAY}}/Main.java
    curl --cookie "session=$SESSION" https://adventofcode.com/{{YEAR}}/day/{{DAY}}/input -o cmd/{{YEAR}}/day_{{DAY}}/input.txt

run YEAR DAY:
    echo "Advent of code - day {{DAY}}"
    if [ -f cmd/{{YEAR}}/day_{{DAY}}/main.go ]; then go run cmd/{{YEAR}}/day_{{DAY}}/main.go; fi
    if [ -f cmd/{{YEAR}}/day_{{DAY}}/Main.java ]; then java cmd/{{YEAR}}/day_{{DAY}}/Main.java; fi

remove YEAR DAY:
    rm -rf cmd/{{YEAR}}/day_{{DAY}}