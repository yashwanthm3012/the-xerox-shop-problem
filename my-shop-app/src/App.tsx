import React, { useState } from 'react';
import { invoke } from '@tauri-apps/api/core';

const App: React.FC = () => {
  const [status, setStatus] = useState<string>('');

  const handleFileUpload = async (event: React.ChangeEvent<HTMLInputElement>) => {
    const file = event.target.files?.[0];
    if (file) {
      setStatus('Uploading...');
      try {
        const result = await invoke<{ status: string }>('upload_file', {
          path: file.name, // Or full path if needed
        });
        setStatus(result.status);
      } catch (error) {
        console.error(error);
        setStatus('Upload failed');
      }
    }
  };

  return (
    <div>
      <h1>Shop Owner - File Upload</h1>
      <input type="file" onChange={handleFileUpload} />
      <p>Status: {status}</p>
    </div>
  );
};

export default App;
