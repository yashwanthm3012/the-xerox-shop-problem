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
    <div className="max-w-md mx-auto mt-10 p-6 bg-white rounded-xl shadow-md space-y-6 font-sans">
      <h2 className="text-2xl font-bold text-gray-800 text-center">Upload PDF for Printing</h2>

      <form onSubmit={handleSubmit} className="space-y-4">
        <div>
          <label className="block text-sm font-medium text-gray-700">Select File</label>
          <input
            type="file"
            onChange={(e) => setFile(e.target.files[0])}
            required
            className="mt-1 block w-full text-sm text-gray-900 file:mr-4 file:py-2 file:px-4 file:rounded-lg file:border-0 file:text-sm file:font-semibold file:bg-blue-50 file:text-blue-700 hover:file:bg-blue-100"
          />
        </div>

        <div>
          <label className="block text-sm font-medium text-gray-700">Print Type</label>
          <select
            value={printType}
            onChange={(e) => setPrintType(e.target.value)}
            className="mt-1 block w-full border border-gray-300 rounded-lg p-2 text-gray-700"
          >
            <option value="bw">Black & White</option>
            <option value="color">Color</option>
            <option value="both">Both</option>
          </select>
        </div>

        {printType === "both" ? (
          <>
            <div>
              <label className="block text-sm font-medium text-gray-700">B/W Pages</label>
              <input
                type="text"
                placeholder="e.g., 1-3,5"
                value={bwPages}
                onChange={(e) => setBwPages(e.target.value)}
                required
                className="mt-1 block w-full border border-gray-300 rounded-lg p-2"
              />
            </div>
            <div>
              <label className="block text-sm font-medium text-gray-700">Color Pages</label>
              <input
                type="text"
                placeholder="e.g., 4,6"
                value={colorPages}
                onChange={(e) => setColorPages(e.target.value)}
                required
                className="mt-1 block w-full border border-gray-300 rounded-lg p-2"
              />
            </div>
          </>
        ) : (
          <div>
            <label className="block text-sm font-medium text-gray-700">Pages</label>
            <input
              type="text"
              placeholder="e.g., 1-5,7"
              value={pages}
              onChange={(e) => setPages(e.target.value)}
              required
              className="mt-1 block w-full border border-gray-300 rounded-lg p-2"
            />
          </div>
        )}

        <button
          type="submit"
          className="w-full bg-blue-600 hover:bg-blue-700 text-white font-medium py-2 px-4 rounded-lg transition"
        >
          Upload
        </button>
      </form>

      {error && (
        <div className="bg-red-100 text-red-700 p-3 rounded-lg text-sm">
          <strong>Error:</strong> {error}
        </div>
      )}

      {response && (
        <div className="bg-green-100 text-green-800 p-3 rounded-lg text-sm">
          <h3 className="font-semibold mb-1">Upload Successful!</h3>
          <pre className="whitespace-pre-wrap">{JSON.stringify(response, null, 2)}</pre>
        </div>
      )}
    </div>
  );
}

export default FileUpload;
