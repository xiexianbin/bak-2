<!DOCTYPE html>

<html>
<head>
  <title>Beego</title>
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
  <link rel="shortcut icon" href="https://www.xiexianbin.cn/favicon.ico" type="image/x-icon" />

  <style type="text/css">
    *,body {
      margin: 0px;
      padding: 0px;
    }

    body {
      margin: 0px;
      font-family: "Helvetica Neue", Helvetica, Arial, sans-serif;
      font-size: 14px;
      line-height: 20px;
      background-color: #fff;
    }

    header,
    .main,
    footer {
      width: 1080px;
      margin-left: auto;
      margin-right: auto;
    }

    .logo {
      background-image: url('https://www.xiexianbin.cn/images/logo/logo-256x256.png');
      background-repeat: no-repeat;
      -webkit-background-size: 80px 80px;
      background-size: 80px 80px;
      background-position: left;
      text-align: center;
      font-size: 42px;
      padding: 40px 0;
      font-weight: normal;
      text-shadow: 0px 1px 2px #ddd;
    }

    header {
      padding: 20px 0;
    }

    footer {
      line-height: 1.8;
      text-align: center;
      padding: 50px 0;
      color: #999;
    }

    .main {
      text-align: left;
      font-size: 16px;
    }

    h4 {
      padding-top: 20px;
    }

    a {
      color: #444;
      text-decoration: none;
    }

    .backdrop {
      position: absolute;
      width: 100%;
      height: 100%;
      box-shadow: inset 0px 0px 100px #ddd;
      z-index: -1;
      top: 0px;
      left: 0px;
    }
  </style>
</head>

<body>
  <header>
    <h2 class="logo">Welcome to Aliyun CDN 404 Page For hugo</h2>
  </header>
  <div class="main">
    <h4>receive https://www.xiexianbin.cn 404 page api call.</h4>

    <h4>Task</h4>
    <ul>
      <li><a target="blank" href="/task?last=7&offset=0">/task?last=7&offset=0</a></li>
      <li><a target="blank" href="/task?last=14&offset=7">/task?last=14&offset=7</a></li>
      <li><a target="blank" href="/task?last=21&offset=14">/task?last=21&offset=14</a></li>
      <li><a target="blank" href="/task?last=28&offset=21">/task?last=28&offset=21</a></li>
      <li><a target="blank" href="/task?last=35&offset=28">/task?last=35&offset=28</a></li>
    </ul>

    <h4>Page404</h4>
    <ul>
      <li><a target="blank" href="/page404">/page404</a></li>
      <li><a target="blank" href="/page404?group=true">/page404?group=true</a></li>
    </ul>

    <h4>Clean All logs in DBs</h4>
    <ul>
      <li>
        <a target="blank" href="/page404?cleandb=true">/page404?cleandb=true</a>
      </li>
    </ul>
  </div>
  <footer>
    <div class="author">
      Official website:
      <a target="blank" href="http://{{.Website}}">{{.Website}}</a> /
      Contact me:
      <a class="email" href="mailto:{{.Email}}">{{.Email}}</a>
    </div>
  </footer>
  <div class="backdrop"></div>

  <script src="/static/js/reload.min.js"></script>
</body>
</html>
