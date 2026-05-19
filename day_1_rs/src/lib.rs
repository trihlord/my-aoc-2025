pub fn password(document: &str) -> usize {
    document
        .lines()
        .map(|s| {
            let mut it = s.chars();
            match (it.next(), it.as_str().parse::<i32>()) {
                (Some('L'), Ok(v)) => -v,
                (Some('R'), Ok(v)) => v,
                _ => panic!(),
            }
        })
        .scan(50, |state, x| {
            *state = (*state + x).rem_euclid(100);
            Some(*state)
        })
        .filter(|x| *x == 0)
        .count()
}

#[cfg(test)]
mod tests {
    use std::fs;

    use super::*;

    #[test]
    fn example() {
        let document =
            fs::read_to_string("tests_data/example.txt").expect("Missing example test data");
        let result = password(&document);
        assert_eq!(result, 3);
    }

    #[test]
    fn input() {
        let document =
            fs::read_to_string("tests_data/input.txt").expect("Missing input test data");
        let result = password(&document);
        assert_eq!(result, 1182)
    }
}
