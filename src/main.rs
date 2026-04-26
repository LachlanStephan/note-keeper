use core::panic;
use serde::Serialize;
use serde_json;
use std::{
    env::home_dir,
    fs::{self, File},
    path::Path,
    time::SystemTime,
};
use time::OffsetDateTime;

static NVIM: &str = "nvim";

#[derive(Debug)]
struct ConfigPaths {
    note_file_path: String,
    config_file_path: String,
    root_dir_path: String,
}

#[derive(Serialize)]
struct Config {
    last_opened: String,
    editor: String,
}

#[derive(Debug)]
struct Application {
    config_paths: ConfigPaths,
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
        },
    };

    if !file_exists(&app.config_paths.root_dir_path) {
        println!("we need to create the stuff");
        scaffold_app(&app);
    }

    let formatted_time = get_curr_time_formatted();
    let last_opened_time = get_last_opened_time();

    if formatted_time != last_opened_time {
        // new day
        write_file(
            &app.config_paths.config_file_path,
            &get_config_values(&formatted_time),
        );
        write_file(&app.config_paths.note_file_path, &formatted_time);
        // create CMD for opening the file in editor of choice
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

// TODO: Implement method
fn read_file_json(file_path: &String) -> String {
    return String::from("value");
}

fn get_curr_time_formatted() -> String {
    let now = SystemTime::now();
    let date_time: OffsetDateTime = now.into();
    return format!(
        "{} {} {} ({})",
        date_time.day(),
        date_time.month(),
        date_time.year(),
        date_time.weekday()
    );
}

// TODO: Implement method
fn get_last_opened_time() -> String {
    // config_data = read_file_json(file_path)
    return String::from("");
}

fn get_config_values(curr_time: &String) -> String {
    let data = Config {
        last_opened: curr_time.to_string(),
        editor: NVIM.to_string(),
    };
    return serde_json::to_string(&data).unwrap();
}
