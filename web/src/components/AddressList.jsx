import React from "react";
import styles from '../App.css';
import Address from "./Address";

const AddressList = ({ addresses }) => {
  if (addresses !== null && addresses.length != 0) {
    return (
      <div>
        {addresses.map(a =>
          <Address address={a} key={a.id} />
        )}
      </div>
    );
  }
  return (
    <div></div>
  )
};

export default AddressList;