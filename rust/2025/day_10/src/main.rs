use std::collections::VecDeque;

fn main() {
    println!("Advent of Code 2025!");
    part1();
    // part2();
}

#[derive(Debug)]
struct Machine {
    lights: Vec<u64>,
    buttons: Vec<Vec<u64>>,
    joltage: Vec<u64>,
}

fn parse_input(input: &str) -> Machine {
    let splitted: Vec<&str> = input.split(" ").collect();

    let lights: Vec<u64> = splitted
        .first()
        .unwrap()
        .trim_matches(|c| c == '[' || c == ']')
        .chars()
        .map(|x| if x == '.' { 0 } else { 1 })
        .collect();

    let buttons: Vec<Vec<u64>> = splitted
        .iter()
        .skip(1)
        .filter(|x| x.contains("(") || x.contains(")"))
        .map(|x| {
            x.trim_matches(|c| c == '(' || c == ')')
                .split(",")
                .map(|n| n.parse::<u64>().unwrap())
                .collect()
        })
        .collect();

    let joltage: Vec<u64> = splitted
        .last()
        .unwrap()
        .trim_matches(|x| x == '{' || x == '}')
        .split(",")
        .map(|x| x.parse::<u64>().unwrap())
        .collect();

    Machine {
        lights,
        buttons,
        joltage,
    }
}

// fn parse_input_v2(input: &str) -> Machine {
//     input.split(' ').fold(
//         Machine {
//             lights: vec![],
//             buttons: vec![],
//             joltage: vec![],
//         },
//         |mut m, t| {
//             match t.as_bytes()[0] {
//                 b'[' => {
//                     m.lights = t
//                         .trim_matches(|c| c == '[' || c == ']')
//                         .chars()
//                         .map(|c| (c == '#') as u64)
//                         .collect()
//                 }
//                 b'(' => m.buttons.push(
//                     t.trim_matches(|c| c == '(' || c == ')')
//                         .split(',')
//                         .map(|n| n.parse().unwrap())
//                         .collect(),
//                 ),
//                 b'{' => {
//                     m.joltage = t
//                         .trim_matches(|c| c == '{' || c == '}')
//                         .split(',')
//                         .map(|n| n.parse().unwrap())
//                         .collect()
//                 }
//                 _ => {}
//             }
//             m
//         },
//     )
// }

struct Node {
    value: u64,
    steps: usize,
    parent: Option<usize>,
    children: Vec<usize>,
}

struct Tree {
    nodes: Vec<Node>,
    expected_output: u64,
    width: usize,
}

impl Tree {
    fn new(root: u64, expected_output: u64, width: usize) -> Self {
        Tree {
            nodes: vec![Node {
                value: root,
                steps: 0,
                parent: None,
                children: vec![],
            }],
            expected_output,
            width,
        }
    }

    fn add_child(&mut self, parent: usize, value: u64) -> usize {
        let idx = self.nodes.len();
        self.nodes.push(Node {
            value,
            steps: self.nodes[parent].steps + 1,
            parent: Some(parent),
            children: vec![],
        });
        self.nodes[parent].children.push(idx);
        idx
    }

    fn build(&mut self, root_index: usize, items: &[u64], start: usize) -> Option<usize> {
        let mut queue: VecDeque<(usize, usize)> = VecDeque::new();
        queue.push_back((root_index, start));

        while let Some((parent_index, start)) = queue.pop_front() {
            for (i, item) in items.iter().enumerate().skip(start) {
                let parent_val = self.nodes[parent_index].value;
                let new_value = parent_val ^ item;

                if new_value == self.expected_output {
                    return Some(self.nodes[parent_index].steps + 1);
                }

                let child_index = self.add_child(parent_index, new_value);
                queue.push_back((child_index, i + 1));
            }
        }

        None
    }
}

fn part1() {
    println!("Part 1");

    let input: &str = include_str!("../input.txt").trim();

    let splitted_input: Vec<&str> = input.split("\n").collect();
    let mut amount: usize = 0;

    for input in splitted_input {
        let machine: Machine = parse_input(input);
        let lights_size = machine.lights.len();

        let expected_output: u64 = machine
            .lights
            .iter()
            .enumerate()
            .fold(0, |acc, (i, &b)| acc | (b << (lights_size - 1 - i)));

        let mut tree = Tree::new(0, expected_output, lights_size);

        let buttons_value: Vec<u64> = machine
            .buttons
            .iter()
            .map(|x| {
                x.iter()
                    .fold(0, |acc, b| acc | (1 << (lights_size as u64 - 1 - b)))
            })
            .collect();

        if let Some(x) = tree.build(0, &buttons_value, 0) {
            amount += x;
        };
    }

    println!("Amount: {}", amount)
}

#[allow(dead_code)]
fn part2() {
    println!("Part 2");
}
