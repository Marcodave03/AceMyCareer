import React, { useState } from "react"
import "./Navbas.scss"
import logo from "../../public/navbar/images/logo.jpg"

function Navbar() {
    var [ isOpen, setIsOpen ] = useState(false)
    let temp = () => {
        setIsOpen(!isOpen)
    }
    return (
        <div className={(isOpen)?"navbar-side active":"navbar-side"}onMouseEnter={temp} onMouseLeave={temp}>
            <div className="navbar-top">

                <div className="navbar-header">
                    <img src={logo} alt="" className="logo" />
                    <h1>RentalTool</h1>
                </div>

                <ul className="navbar-items">

                    <li className="navbar-item">
                        <i className="icon"></i>
                        <h1 className="title">Profile</h1>
                    </li>

                    <li className="navbar-item">
                        <i className="icon"></i>
                        <h1 className="title">Home</h1>
                    </li>

                    <li className="navbar-item">
                        <i className="icon"></i>
                       <h1 className="title">Shop</h1>
                    </li>

                    <li className="navbar-item">
                        <i className="icon"></i>
                        <h1 className="title">Cart</h1>
                    </li>


                </ul>

            </div>
            <div className="navbar-bottom">

                <ul>
                    <li className="navbar-item">
                        <i className="icon"></i>
                        <h1 className="title">Logout</h1>
                    </li>
                </ul>
            </div>
        </div>
    )
}

export default Navbar

