import React from "react";
import FullRecord from "./FullRecord";
import styles from "../App.css";

const FullRecordList = ({ recs }) => {
  if (recs === null || recs === undefined || recs.length === 0) {
    return (<div>Записей не найдено</div>)
  }
  return (
    <div>
      {recs.map(rec =>
        <FullRecord rec={rec} key={rec.address.id} />
      )}
    </div>
  );
};

export default FullRecordList;