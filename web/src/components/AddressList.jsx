import React from "react";
import Address from "./Address";

const AddressList = ({ addresses }) => {
  if (addresses === null || addresses === undefined || addresses.length == 0) {
    return (
      <div>Записей не найдено</div>
    )
  }
  return (
    <div>
      {addresses.map(a =>
        <Address address={a} key={a.id} />
      )}
    </div>
  );
};

export default AddressList;