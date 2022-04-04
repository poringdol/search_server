import React from "react";
import '../App.css';
import Person from "./Person";

const PersonList = ({ persons }) => {
  if (persons === null || persons == undefined || persons.length == 0) {
    return (<div>Записей не найдено</div>)
  }
  return (
    <div>
      {persons.map(p =>
        <Person person={p} key={p.phone} />
      )}
    </div>
  );
};


export default PersonList