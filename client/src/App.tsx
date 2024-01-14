import React from "react";
import CsvUpload from "./components/csv-upload/csv-upload";

const handleFileUpload = (file: File) => {
  // Do something with the uploaded file
  console.log(file);
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
