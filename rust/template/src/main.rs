use std::time::Instant;

use aoc___YEAR___day___DAY__::{part1, part2};

fn main() {
    println!("Advent of Code __YEAR__ - day __DAY__");
    let input = include_str!("../input.txt").trim_end_matches('\n');

    let t = Instant::now();
    let p1 = part1(input);
    println!("Part 1: {p1}  ({:?})", t.elapsed());

    let t = Instant::now();
    let p2 = part2(input);
    println!("Part 2: {p2}  ({:?})", t.elapsed());
}
