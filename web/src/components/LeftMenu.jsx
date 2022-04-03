import React from "react";

const LeftMenu = () => {
    return (
        <div style={{ position: 'absolute', paddingLeft: '40px' }}>
            <button>Найти по телефону</button> <br />
            <button>Найти по имени</button> <br />
            <button>Найти по адресу</button> <br />
        </div>
    );
}

export default LeftMenu