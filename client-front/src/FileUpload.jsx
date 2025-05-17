import React, { useState, useEffect } from "react";

function FileUpload() {
  const [file, setFile] = useState(null);
  const [fileURL, setFileURL] = useState(null); // For preview
  const [printType, setPrintType] = useState("bw");
  const [bwPages, setBwPages] = useState("");
  const [colorPages, setColorPages] = useState("");
  const [pages, setPages] = useState("");
  const [response, setResponse] = useState(null);
  const [error, setError] = useState("");

  useEffect(() => {
    if (file) {
      const url = URL.createObjectURL(file);
      setFileURL(url);
      return () => URL.revokeObjectURL(url); // Cleanup
    }
    setFileURL(null);
  }, [file]);

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
    <div className="max-w-md mx-auto mt-10 p-6 bg-gray-200 border border-gray-500 shadow-[4px_4px_0px_0px_rgba(0,0,0,0.8)] font-mono text-sm space-y-6">
      <h2 className="text-xl font-bold text-black text-center border-b border-gray-500 pb-2">Print Portal</h2>

      <form onSubmit={handleSubmit} className="space-y-4">
        <div>
          <label className="block text-xs font-bold text-gray-800 mb-1">Select File</label>
          <input
            type="file"
            accept=".pdf"
            onChange={(e) => setFile(e.target.files[0])}
            className="w-full text-xs p-1 bg-white border border-gray-400 shadow-inner"
          />
        </div>

        {fileURL && file?.type === "application/pdf" && (
          <div>
            <label className="block text-xs font-bold text-gray-800 mb-1">Preview</label>
            <div className="border border-gray-400 shadow-inner bg-white h-64 overflow-hidden">
              <iframe
                src={fileURL}
                className="w-full h-full"
                title="PDF Preview"
              ></iframe>
            </div>
          </div>
        )}

        <div>
          <label className="block text-xs font-bold text-gray-800 mb-1">Print Type</label>
          <select
            value={printType}
            onChange={(e) => setPrintType(e.target.value)}
            className="w-full text-xs p-1 bg-white border border-gray-400 shadow-inner"
          >
            <option value="bw">Black & White</option>
            <option value="color">Color</option>
            <option value="both">Both</option>
          </select>
        </div>

        {printType === "both" ? (
          <>
            <div>
              <label className="block text-xs font-bold text-gray-800 mb-1">B/W Pages</label>
              <input
                type="text"
                placeholder="e.g., 1-3,5"
                value={bwPages}
                onChange={(e) => setBwPages(e.target.value)}
                className="w-full text-xs p-1 bg-white border border-gray-400 shadow-inner"
              />
            </div>
            <div>
              <label className="block text-xs font-bold text-gray-800 mb-1">Color Pages</label>
              <input
                type="text"
                placeholder="e.g., 4,6"
                value={colorPages}
                onChange={(e) => setColorPages(e.target.value)}
                className="w-full text-xs p-1 bg-white border border-gray-400 shadow-inner"
              />
            </div>
          </>
        ) : (
          <div>
            <label className="block text-xs font-bold text-gray-800 mb-1">Pages</label>
            <input
              type="text"
              placeholder="e.g., 1-5,7"
              value={pages}
              onChange={(e) => setPages(e.target.value)}
              className="w-full text-xs p-1 bg-white border border-gray-400 shadow-inner"
            />
          </div>
        )}

        <button
          type="submit"
          className="w-full bg-gray-300 text-black font-bold py-2 border border-black shadow-[2px_2px_0px_rgba(0,0,0,0.8)] active:translate-x-[2px] active:translate-y-[2px] active:shadow-none"
        >
          ▶ Upload
        </button>
      </form>

      {error && (
        <div className="bg-red-100 text-red-700 p-2 border border-red-400 shadow-inner text-xs">
          <strong>Error:</strong> {error}
        </div>
      )}

      {response && (
        <div className="bg-green-100 text-green-800 p-2 border border-green-400 shadow-inner text-xs">
          <h3 className="font-bold mb-1">Upload Successful!</h3>
          <pre className="whitespace-pre-wrap">{JSON.stringify(response, null, 2)}</pre>
        </div>
      )}
    </div>
  );
}

export default FileUpload;
