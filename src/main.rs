use core::panic;
use serde::{Deserialize, Serialize};
use serde_json;
use std::{
    env::home_dir,
    fs::{self, File},
    io::Result,
    io::Write,
    path::Path,
    process::Command,
    time::SystemTime,
};
use time::OffsetDateTime;

const NVIM: &str = "nvim";

#[derive(Debug)]
struct ConfigPaths {
    note_file_path: String,
    config_file_path: String,
    root_dir_path: String,
}

#[derive(Serialize, Deserialize)]
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
        scaffold_app(&app);
    }

    handle_if_new_day(&app);
    open_note(&app.config_paths.note_file_path);
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
    write_file(&app.config_paths.note_file_path, "# Note\n\n");
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

fn handle_if_new_day(app: &Application) {
    let formatted_time = get_curr_time_formatted();
    let last_opened_time = get_last_opened_time(&app.config_paths.config_file_path);

    if formatted_time == last_opened_time {
        return;
    }

    write_file(
        &app.config_paths.config_file_path,
        &get_config_values(&formatted_time),
    );

    let result = append_file(&app.config_paths.note_file_path, &formatted_time);
    if !result.is_ok() {
        panic!("could not append new day heading");
    }
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

fn get_last_opened_time(file_path: &String) -> String {
    let config_data = read_file_json(file_path);
    if config_data == "" {
        return String::from("");
    }

    let config: Config =
        serde_json::from_str(&config_data).expect("could not convert config to struct");

    return config.last_opened;
}

fn get_config_values(curr_time: &String) -> String {
    let data = Config {
        last_opened: curr_time.to_string(),
        editor: NVIM.to_string(),
    };
    return serde_json::to_string(&data).unwrap();
}

fn append_file(file_path: &String, contents: &str) -> Result<()> {
    let mut file = File::options().append(true).open(file_path)?;
    writeln!(&mut file, "## {}", contents)?;
    Ok(())
}

fn read_file_json(file_path: &String) -> String {
    let err = "could not read the config";
    let result = fs::read_to_string(file_path).expect(err);
    return result;
}

fn open_note(note_path: &String) {
    Command::new(NVIM.to_string())
        .args([note_path])
        .status()
        .expect("failed to open note file");
}