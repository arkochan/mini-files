async function uploadFile() {
  const fileInput = document.getElementById('fileInput');
  const file = fileInput.files[0];
  if (!file) {
    alert('Please select a file');
    return;
  }

  try {
    // Get presigned URL for upload
    const response = await fetch(`/get-upload-url?filename=${file.name}`);
    const presignedUrl = await response.text();

    // Upload file using presigned URL
    await fetch(presignedUrl, {
      method: 'PUT',
      body: file
    });

    alert('File uploaded successfully!');
  } catch (error) {
    console.error('Error:', error);
    alert('Upload failed');
  }
}

async function downloadFile() {
  const filename = document.getElementById('filename').value;
  if (!filename) {
    alert('Please enter a filename');
    return;
  }

  try {
    const response = await fetch(`/get-download-url?filename=${filename}`);
    const presignedUrl = await response.text();
    // Open download link in new tab
    window.open(presignedUrl, '_blank');
  } catch (error) {
    console.error('Error:', error);
    alert('Failed to get download link');
  }
}

