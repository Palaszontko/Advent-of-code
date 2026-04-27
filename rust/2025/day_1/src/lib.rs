#[derive(Debug)]
pub enum Direction {
    Left,
    Right,
}

#[derive(Debug)]
pub struct Move {
    direction: Direction,
    amount: i32,
}

fn parse(input: &str) -> Vec<Move> {
    input
        .split('\n')
        .map(|x| x.split_at(1))
        .map(|(d, a)| Move {
            direction: match d {
                "L" => Direction::Left,
                "R" => Direction::Right,
                _ => panic!("Invalid"),
            },
            amount: a.trim().parse().unwrap(),
        })
        .collect()
}

pub fn part1(input: &str) -> i64 {
    let data = parse(input);

    let mut start: i32 = 50;
    let mut amount: usize = 0;

    for rotation in data {
        start = match rotation.direction {
            Direction::Left => modular_sub_100(start, rotation.amount),
            Direction::Right => (start + rotation.amount) % 100,
        };
        if start == 0 {
            amount += 1
        }
    }

    amount as i64
}

fn modular_sub_100(a: i32, b: i32) -> i32 {
    ((a - b) % 100 + 100) % 100
}

pub fn part2(input: &str) -> i64 {
    let data = parse(input);

    let mut start: i32 = 50;
    let mut amount: u64 = 0;

    for rotation in data {
        match rotation.direction {
            Direction::Left => {
                let (next_pos, matched) = count_zeros_and_last_pos(start, -rotation.amount);
                start = next_pos;
                amount += matched as u64;
            }
            Direction::Right => {
                let (next_pos, matched) = count_zeros_and_last_pos(start, rotation.amount);
                start = next_pos;
                amount += matched as u64;
            }
        }
    }

    amount as i64
}

fn count_zeros_and_last_pos(start: i32, amount: i32) -> (i32, i32) {
    let end = start + amount;
    let zeros = if amount >= 0 {
        end.div_euclid(100) - start.div_euclid(100)
    } else {
        (start - 1).div_euclid(100) - (end - 1).div_euclid(100)
    };
    (end.rem_euclid(100), zeros)
}
