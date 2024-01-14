import React from "react";
import CsvUpload from "./components/csv-upload/csv-upload";
import axios from "axios";

const handleFileUpload = (file: File) => {
  console.log("got the file!");
  axios
    .get("http://localhost:8080/api/health")
    .then((res) => console.log(res.data))
    .catch((err) => console.error(err));
  const formData = new FormData();
  formData.append("file", file);
  axios
    .post("http://localhost:8080/api/upload", formData, {
      headers: {
        "Content-Type": "multipart/form-data",
      },
    })
    .then((res) => console.log(res.data))
    .catch((err) => console.error(err));
};

const App: React.FC = () => {
  return (
    <div>
      <h1>File Picker Example</h1>
      <CsvUpload onFileUpload={handleFileUpload} />
    </div>
  );
};

export default App;
