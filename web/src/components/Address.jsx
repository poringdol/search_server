import React from "react";
import styles from '../App.css';
import FullRecordList from "./FullRecordList";
import capitalize from "../capitalize"

const addressess = ({address}) => {
  if (address !== null && address.length != 0) {
    return (
      <div>
        <div style={{marginBottom: '10px'}}>
          <h3>АДРЕС</h3>
          <strong>Город: </strong> {capitalize(address.city)}, <strong>улица:</strong> {capitalize(address.street)}, <strong>дом:</strong> {address.house}, <strong>квартира:</strong> {address.office}
        </div>
        <FullRecordList recs={address.full_records} />
      </div>
    );
  }
  return (
    <div></div>
  )
};

export default addressess;