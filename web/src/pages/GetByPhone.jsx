import React, { useState } from "react";
import AddressList from "../components/AddressList";
import Person from "../components/Person";
import MainPageSelect from "../components/MainPageSelect";

const GetByPhone = function () {
  const [phone, setPhone] = useState("")
  const [person, setPerson] = useState(null)
  const [addresses, setAddresses] = useState([])

  function findByPhone() {
    fetch("http://localhost:9000/find_by_phone?phone=" + phone, {
      method: 'GET',
      headers: {
        "Access-Control-Allow-Origin": "*"
      }
    })
      .then(response => response.json())
      .then(result => {
        if (result.error.code === 404) {
          setPerson("")
          setAddresses([])
          return
        }
          setPerson(result.person)
          setAddresses(result.addresses)
      })
    // .catch(e => {
    //   console.log(e);
    //   setData(e)});
  }

  return (
    <div className="content">
      <input className="input" type="text" placeholder="Телефон" value={phone} onChange={e => setPhone(e.target.value)} />
      <button className="input" onClick={findByPhone}>Отправить</button>

      <Person person={person} />

      <AddressList addresses={addresses} />

    </div>
  )
}

export default GetByPhone;