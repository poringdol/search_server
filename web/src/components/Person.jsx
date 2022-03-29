import React from "react";
import '../App.css';
import capitalize from "../capitalize"
import styles from "../App.css";

const Person = ({ person }) => {
  if (person !== null) {
    return (
      <div style={{
        border: '3px solid rgba(0, 0, 150, 0.5)',
        margin: '10px',
        textAlign: 'center',
        padding: '10px'
      }}>
        <strong>Имя</strong>: {capitalize(person.full_name)} <br />
        <strong>email: </strong> {person.email} <br />
        <strong>телефон:</strong> {person.phone} <br />
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