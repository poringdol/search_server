import React, { useState } from "react";
import AddressList from "../components/AddressList";
import Person from "../components/Person";

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
    <div className="input_div">
      <input className="input_text" type="text" placeholder="Телефон" value={phone} onChange={e => setPhone(e.target.value)} />
      <button className="input_text" onClick={findByPhone}>Отправить</button>

      <Person person={person} />

      <AddressList addresses={addresses} />

    </div>
  )
}

export default GetByPhone;