use std::env::home_dir;

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
    let note = root.clone() + "/" + "nk-note.md";
    let config = root.clone() + "/" + "nk-config.txt";

    let app = Application {
        config_paths: ConfigPaths {
            root_dir_path: root,
            note_file_path: note,
            config_file_path: config,
        }
    };

    println!("{:?}", app);
}

fn get_home_dir() -> String {
    let home_dir = home_dir();
    let path = match home_dir {
        Some(path) => path.display().to_string(),
        None => panic!("could not get home dir"),
    };

    return path + "/" + "note-keeper"; // append app root dir
}
