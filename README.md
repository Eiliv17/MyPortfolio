<a name="readme-top"></a>
<!-- Head section -->
<div align="center">
  <a href="https://github.com/Eiliv17/CloudStorageWebApp">
    <img src="./README/Product.svg" alt="Logo">
  </a>

  <h1 align="center">My Portfolio Project</h1>

  <p align="center">
    Portfolio website create in Vue.js with a backend written in Go using the Gin web framework and MongoDB as database
  </p>
</div>

<!-- Table of contents -->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li><a href="#license">License</a></li>
  </ol>
</details>

<!-- About the project section -->
## About The Project
This is the first time I've built a fully fledge portflio, at first I wanted to only use Go's HTML templates so that it would be also indexable by search engines but it's just a too long of a process to do all that and I wanted to try Vue.js since for now I've only built basic HTML, CSS websites with a little bit of JavaScript.

The projects is still a work in progress but most of the functionality is done, I would also like in the future to add Go HTML templates to differenciate between normal user requests and search engine crawlers so that I can serve different pages for the crawlers and the final users, so that the website is still indexed by search engines and still functional to the end user. 

The CSS needs a little bit of cleaning and maybe I need to use the BEM convention to make the code more pleasing to read since for now I've only done scoped styles thanks to Vue.js components.

<p align="right">(<a href="#readme-top">back to top</a>)</p>


<!-- Built with section -->
### Built With

This project was built with the following technologies:

- ![HTML](https://img.shields.io/badge/HTML5-E34F26?style=for-the-badge&logo=html5&logoColor=white)
- ![CSS](https://img.shields.io/badge/CSS3-1572B6?style=for-the-badge&logo=css3&logoColor=white)
- ![JavaScript](https://img.shields.io/badge/javascript-%23323330.svg?style=for-the-badge&logo=javascript&logoColor=%23F7DF1E)
- ![Vue.js](https://img.shields.io/badge/vuejs-%2335495e.svg?style=for-the-badge&logo=vuedotjs&logoColor=%234FC08D)
- ![Golang](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)
- ![MongoDB](https://img.shields.io/badge/MongoDB-4EA94B?style=for-the-badge&logo=mongodb&logoColor=white)

To be more clear the projects uses the Vue.js framework on the front-end with Vue Router to serve dinamically different pages, some page content is static while other like blog posts, projects and the contact form are dinamic, meaning you can interact with those and they change, as for the backend, it was build with the Go language using the Gin framework, I didn't use a Go server for the API and a Node.js/Express server for serving the website since I didn't need to render the website on the server and then idrate the front-end, also it's much faster a Golang server than an Express server. And as a database I'm using MongoDB to store the blog posts, projects and the contact form data.

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- license section -->
## License
[![License: GPL v3](https://img.shields.io/badge/License-GPLv3-blue.svg?style=for-the-badge&logoColor=white)](https://www.gnu.org/licenses/gpl-3.0)

Distributed under the GPL v3 License. See [LICENSE](LICENSE) for more information.

<p align="right">(<a href="#readme-top">back to top</a>)</p>