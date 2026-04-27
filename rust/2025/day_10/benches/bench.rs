use aoc_2025_day_10 as solution;
use divan::{AllocProfiler, Bencher, counter::BytesCount};

#[global_allocator]
static ALLOC: AllocProfiler = AllocProfiler::system();

fn main() {
    divan::main();
}

#[divan::bench(sample_count = 200, min_time = 0.5)]
fn part1(bencher: Bencher) {
    let input = include_str!("../input.txt").trim_end_matches('\n');
    bencher
        .counter(BytesCount::of_str(input))
        .bench(|| solution::part1(divan::black_box(input)));
}

#[divan::bench(sample_count = 200, min_time = 0.5)]
fn part2(bencher: Bencher) {
    let input = include_str!("../input.txt").trim_end_matches('\n');
    bencher
        .counter(BytesCount::of_str(input))
        .bench(|| solution::part2(divan::black_box(input)));
}
