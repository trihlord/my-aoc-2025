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
        .sum()
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        let result = password(
            r#"L68
L30
R48
L5
R60
L55
L1
L99
R14
L82"#,
        );
        assert_eq!(result, 3);
    }
}
