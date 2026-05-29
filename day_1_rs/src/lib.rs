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
        let document = fs::read_to_string("tests_data/input.txt").expect("Missing input test data");
        let result = password(&document);
        assert_eq!(result, 1182)
    }
}

pub mod v2 {
    pub fn password(document: &str) -> i32 {
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
            .fold((50, 0), |acc, x| {
                (
                    (acc.0 + x).rem_euclid(100),
                    if acc.0 == 0 {
                        acc.1
                    } else {
                        acc.1
                            + (acc.0 + x).div_euclid(100).abs()
                            + if acc.0 + x == 0 { 1 } else { 0 }
                    },
                )
            })
            .1
    }

    #[cfg(test)]
    mod test {
        use std::fs;

        use super::*;

        #[test]
        fn example() {
            let document =
                fs::read_to_string("tests_data/example.txt").expect("Missing example test data");
            let result = password(&document);
            assert_eq!(result, 6);
        }

        #[test]
        fn example_1() {
            let document =
                fs::read_to_string("tests_data/example_1.txt").expect("Missing example test data");
            let result = password(&document);
            assert_eq!(result, 1);
        }

        #[test]
        fn example_2() {
            let document =
                fs::read_to_string("tests_data/example_2.txt").expect("Missing example test data");
            let result = password(&document);
            assert_eq!(result, 2);
        }

        #[test]
        fn example_3() {
            let document =
                fs::read_to_string("tests_data/example_3.txt").expect("Missing example test data");
            let result = password(&document);
            assert_eq!(result, 3);
        }

        #[test]
        fn example_4() {
            let document =
                fs::read_to_string("tests_data/example_4.txt").expect("Missing example test data");
            let result = password(&document);
            assert_eq!(result, 4);
        }

        #[test]
        fn example_5() {
            let document =
                fs::read_to_string("tests_data/example_5.txt").expect("Missing example test data");
            let result = password(&document);
            assert_eq!(result, 5);
        }

        #[test]
        fn input() {
            let document =
                fs::read_to_string("tests_data/input.txt").expect("Missing input test data");
            let result = password(&document);
            assert_eq!(result, 6907)
        }
    }
}
