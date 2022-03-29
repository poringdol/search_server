import React from "react";
import '../App.css';
import capitalize from "../capitalize"
import styles from "../App.css";

const FullRecord = ({rec}) => {
    return (
        <div style={{ border: '1px solid rgba(0, 0, 100, 0.5)', padding: '10px' }}>
            <strong>Имя: </strong> {capitalize(rec.person.full_name)}, <strong>телефон:</strong> {rec.person.phone}, <strong>email: </strong> {rec.person.email} <br />
            {/* <strong>Город: </strong> {rec.address.city}, <strong>улица:</strong> {rec.address.street}, <strong>дом:</strong> {rec.address.house}, 
            <strong>подъезд:</strong>, {rec.address.entrance}, <strong>этаж:</strong> {rec.address.floor}, <strong>квартира:</strong> {rec.address.office},
            <strong>домофон:</strong> {rec.address.door_code} */}
            <strong>сумма заказа: </strong> {rec.amount_charged}, <strong>дата заказа: </strong> {rec.created_at} <br />
            <strong>комментарий: </strong> {rec.address_comment}, <br />
        </div>
    );
};

export default FullRecord;