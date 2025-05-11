use tauri::{command, generate_handler};

#[tauri::command]
fn upload_file(file_data: Vec<u8>, filename: String) -> Result<UploadResult, String> {
    // Example of saving the file to the disk
    use std::fs::File;
    use std::io::Write;
    let file_path = format!("./uploads/{}", filename);
    let mut file = match File::create(&file_path) {
        Ok(file) => file,
        Err(_) => return Err("Failed to create file".to_string()),
    };

    match file.write_all(&file_data) {
        Ok(_) => Ok(UploadResult {
            status: "File uploaded successfully".to_string(),
        }),
        Err(_) => Err("Failed to write file".to_string()),
    }
}

#[derive(serde::Serialize)]
struct UploadResult {
    status: String,
}

fn main() {
    tauri::Builder::default()
        .invoke_handler(generate_handler![upload_file])
        .run(tauri::generate_context!())
        .expect("error while running tauri application");
}
