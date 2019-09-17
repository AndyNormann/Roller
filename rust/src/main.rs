extern crate rand;

use regex::Regex;
use std::env;
use rand::Rng;

fn main() {
    let args: Vec<String> = env::args().collect();
    let re = Regex::new(r"(\d*)?d(\d*)\+?(\d*)?").unwrap();

    if args.len() != 2 {
        println!("Usage: ./{} <roll>", args[0]);
        return;
    }

    let groups = re.captures(&args[1]).unwrap();

    let dice   = groups.get(2).map_or(0, |v| v.as_str().parse::<i32>().unwrap());
    let amount = groups.get(1).map_or(1, |v| v.as_str().parse::<i32>().unwrap());
    let plus   = groups.get(3).map_or(0, |v| v.as_str().parse::<i32>().unwrap());

    let mut result = plus;
    let mut result_vector = vec!();

    let mut rng = rand::thread_rng();

    for _ in 0..amount {
        let cur = rng.gen_range(1, dice);
        result += cur;
        result_vector.push(cur.to_string());
    }

    println!("{} +{} = {}", result_vector.join(" "), plus, result);
}
