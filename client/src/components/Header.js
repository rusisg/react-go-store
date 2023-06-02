import React from 'react'

export default function Header() {
  return (
      <header>
          <div className='header_container'>
                <a href="/" className="logo" images="">
                    <span>Eurostyle</span>
                </a>
               <nav>
                <ul class="nav_links">
                    <li><a class="homeButton" href="/">Home</a></li>
                    <li><a class="catalogButton" href="/catalog">Catalogs</a></li>
                    <li><a class="contactButton" href="/contact">Contacts</a></li>
                      <a href="/order">

                    </a>
                </ul>
            </nav>
        </div>
    </header>
  )
}
