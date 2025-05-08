import React, { useState } from "react";

function FileUpload() {
  const [file, setFile] = useState(null);
  const [printType, setPrintType] = useState("bw");
  const [bwPages, setBwPages] = useState("");
  const [colorPages, setColorPages] = useState("");
  const [pages, setPages] = useState("");
  const [response, setResponse] = useState(null);
  const [error, setError] = useState("");

  const handleSubmit = async (e) => {
    e.preventDefault();

    if (!file) {
      setError("Please select a file.");
      return;
    }

    const formData = new FormData();
    formData.append("file", file);
    formData.append("printType", printType);

    if (printType === "both") {
      formData.append("bwPages", bwPages);
      formData.append("colorPages", colorPages);
    } else {
      formData.append("pages", pages);
    }

    try {
      const res = await fetch("http://localhost:3000/upload", {
        method: "POST",
        body: formData,
      });

      const text = await res.text();

      if (!res.ok) {
        setError(text);
        setResponse(null);
        return;
      }

      const data = JSON.parse(text);
      setResponse(data);
      setError("");
    } catch (err) {
      console.error("Network error:", err);
      setError("Failed to connect to server.");
      setResponse(null);
    }
  };

  return (
    <div style={{ padding: "20px", fontFamily: "Arial" }}>
      <h2>Upload PDF for Printing</h2>

      <form onSubmit={handleSubmit} style={{ marginBottom: "20px" }}>
        <div>
          <input type="file" onChange={(e) => setFile(e.target.files[0])} required />
        </div>

        <div>
          <label>Print Type: </label>
          <select value={printType} onChange={(e) => setPrintType(e.target.value)}>
            <option value="bw">Black & White</option>
            <option value="color">Color</option>
            <option value="both">Both</option>
          </select>
        </div>

        {printType === "both" ? (
          <>
            <div>
              <input
                type="text"
                placeholder="B/W Pages (e.g., 1-3,5)"
                value={bwPages}
                onChange={(e) => setBwPages(e.target.value)}
                required
              />
            </div>
            <div>
              <input
                type="text"
                placeholder="Color Pages (e.g., 4,6)"
                value={colorPages}
                onChange={(e) => setColorPages(e.target.value)}
                required
              />
            </div>
          </>
        ) : (
          <div>
            <input
              type="text"
              placeholder="Pages (e.g., 1-5,7)"
              value={pages}
              onChange={(e) => setPages(e.target.value)}
              required
            />
          </div>
        )}

        <button type="submit">Upload</button>
      </form>

      {error && <div style={{ color: "red" }}>Error: {error}</div>}

      {response && (
        <div style={{ color: "green" }}>
          <h3>Upload Successful!</h3>
          <pre>{JSON.stringify(response, null, 2)}</pre>
        </div>
      )}
    </div>
  );
}

export default FileUpload;
