use std::fs;

pub fn read_file(path: &str) -> String {
    let contents = fs::read_to_string(path)
        .unwrap_or_else(|e| panic!("failed to read file {path:?}: {e}"));
    contents
        .strip_suffix('\n')
        .map(str::to_owned)
        .unwrap_or(contents)
}
