import React, { Component } from 'react'
import Select from 'react-select'

const options = [
    { value: 'byPhone', label: 'По телефону' },
    { value: 'byName', label: 'По имени' },
    { value: 'byAddress', label: 'По адресу' }
  ]  

const MainPageSelect = () => {
    <Select options={options} />
}

export default MainPageSelect