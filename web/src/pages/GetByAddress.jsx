import React, {useState} from "react";
import FullRecordList from "../components/FullRecordList";

const GetByAddress = function () {
  const [city, setCity] = useState("")
  const [street, setStreet] = useState("")
  const [house, setHouse] = useState("")
  const [office, setOffice] = useState("")
  const [addresses, setAddresses] = useState([])

  function findByAddress() {
    fetch("http://localhost:9000/find_by_address?city="+city+"&street="+street+"&house="+house+"&office="+office, {
      method: 'GET',
      headers: {
        "Access-Control-Allow-Origin": "*"
      }
    })
      .then(response => response.json())
      .then(result => {
        if (result.error.code === 404) {
          setAddresses([])
          return
        }
        setAddresses(result.addresses)
      })
    // .catch(e => {
    //   console.log(e);
    //   setData(e)});
  }

  return (
      <div>
        <input className="input_text" type="text" placeholder="Город" value={city} onChange={e => setCity(e.target.value)} />
        <input className="input_text" type="text" placeholder="Улица" value={street} onChange={e => setStreet(e.target.value)} />
        <input className="input_text" type="text" placeholder="Дом" value={house} onChange={e => setHouse(e.target.value)} />
        <input className="input_text" type="text" placeholder="Квартира" value={office} onChange={e => setOffice(e.target.value)} />
        <button className="input_text" onClick={findByAddress}>Отправить</button>

        <FullRecordList recs={addresses} />
      </div>
  )
}

export default GetByAddress