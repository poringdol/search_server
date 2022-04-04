import React from "react";
import '../App.css';
import capitalize from "../capitalize"

const Person = ({ person }) => {
  if (person !== null) {
    return (
      <div className="person">
        <strong>Имя</strong>: {capitalize(person.full_name)} <br />
        <strong>Телефон:</strong> {person.phone} <br />
        <strong>Email: </strong> {person.email} <br />
      </div>
    );
  }
  return (
    <div>
      <strong>Записей не найдено</strong>
    </div>
  );
};

export default Person