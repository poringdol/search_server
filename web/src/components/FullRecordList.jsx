import React from "react";
import FullRecord from "./FullRecord";
import styles from "../App.css";

const FullRecordList = ({ recs }) => {
  return (
    <div>
      {recs.map(rec =>
        <FullRecord rec={rec} key={rec.address.id} />
      )}
    </div>
  );
};

export default FullRecordList;