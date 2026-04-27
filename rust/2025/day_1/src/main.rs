use std::time::Instant;

use aoc_2025_day_1::{part1, part2};

fn main() {
    println!("Advent of Code 2025 - day 1");
    let input = include_str!("../input.txt").trim_end_matches('\n');

    let t = Instant::now();
    let p1 = part1(input);
    println!("Part 1: {p1}  ({:?})", t.elapsed());

    let t = Instant::now();
    let p2 = part2(input);
    println!("Part 2: {p2}  ({:?})", t.elapsed());
}
