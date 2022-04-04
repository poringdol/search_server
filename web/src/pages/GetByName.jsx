import React, {useState} from "react";
import PersonList from "../components/PersonList";

const GetByName = function () {
    const [surname, setSurname] = useState("")
    const [name, setName] = useState("")
    const [patronymic, setPatronymic] = useState("")
    const [persons, setPersons] = useState([])

    const findByName = () => {
    fetch("http://localhost:9000/find_by_name?surname="+surname+"&name="+name+"&patronymic="+patronymic, {
        method: 'GET',
        headers: {
            "Access-Control-Allow-Origin": "*"
        }
    })
      .then(response => response.json())
      .then(result => {
          if (result.error.code === 404) {
            setPersons([])
              return
          }
        setPersons(result.persons)
        console.log(persons)
      })
    // .catch(e => {
    //   console.log(e);
    //   setData(e)});
}

    return (
      <div>
          <input className="input_text" type="text" placeholder="Фамилия" value={surname} onChange={e => setSurname(e.target.value)} />
          <input className="input_text" type="text" placeholder="Имя" value={name} onChange={e => setName(e.target.value)} />
          <input className="input_text" type="text" placeholder="Отчество" value={patronymic} onChange={e => setPatronymic(e.target.value)} />
          <button className="input_text" onClick={findByName}>Отправить</button>

          <PersonList persons={persons} />
      </div>
    )
}


export default GetByName