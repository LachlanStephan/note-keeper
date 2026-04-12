use core::panic;
use std::{env::home_dir, fs::{self, File}, path::Path};

#[derive(Debug)]
struct ConfigPaths {
    note_file_path: String,
    config_file_path: String,
    root_dir_path: String
}

#[derive(Debug)]
struct Application {
    config_paths: ConfigPaths
}

fn main() {
    let root = String::from(get_home_dir());
    let note = root.clone() + "/" + "note.md";
    let config = root.clone() + "/" + "config.json";

    let app = Application {
        config_paths: ConfigPaths {
            root_dir_path: root,
            note_file_path: note,
            config_file_path: config,
        }
    };

    if !file_exists(&app.config_paths.root_dir_path) {
        println!("we need to create the stuff");
        scaffold_app(&app);
    }
}

fn get_home_dir() -> String {
    let home_dir = home_dir();
    let path = match home_dir {
        Some(path) => path.display().to_string(),
        None => panic!("could not get home dir"),
    };

    return path + "/" + "note-keeper"; // append app root dir
}

fn file_exists(file_path: &String) -> bool {
    return Path::new(file_path).exists();
}

fn scaffold_app(app: &Application) {
    create_dir(&app.config_paths.root_dir_path);
    create_file(&app.config_paths.config_file_path);
    create_file(&app.config_paths.note_file_path);
    write_file(&app.config_paths.note_file_path, "# Note");
}

fn create_dir(dir_path: &String) {
    let result = fs::create_dir(dir_path);
    if !result.is_ok() {
        panic!("could not create dir: {}", dir_path);
    }
}

fn create_file(file_path: &String) {
    let result = File::create(file_path);
    if !result.is_ok() {
        panic!("could create file: {}", file_path);
    }
}

fn write_file(file_path: &String, contents: &str) {
    let result = fs::write(file_path, contents);
    if !result.is_ok() {
        panic!("could not write to file: {}", file_path);
    }
}