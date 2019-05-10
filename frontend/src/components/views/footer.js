import React from 'react';
import { Link } from 'react-router';
import vk from '../../styles/icons/vk-icon-circle.png'


function Footer(props) {
  return (
    <footer>
      <div className="footer">
        <div className="navigation">
          <div className="logo">
            <div className="Logo-img"> </div>
          </div>

          <div className="copyright">
            <p>© 2019 «Культурный город», по вопросам пишите по адресу spb-cc@gmail.com или через форму на сайте</p>
          </div>
            <div className="vk-icon"><Link to='/about'><img src={vk} alt="vk-icon" /></Link></div>

        </div>
      </div>
    </footer>
  )
}

export default Footer
