# 手机UI优化

# 分享优化
```
sed -i "s|share-dialog-side{display:-ms-flexbox;display:flex;|share-dialog-side{display:none;|g" /data/seafile-server-7.0.5/seahub/media/assets/frontend/css/app.b34c01f55f48.css
sed -i "s|width:35rem;height:20rem|width:100%;height:20rem|g" /data/seafile-server-7.0.5/seahub/media/assets/frontend/css/app.b34c01f55f48.css
```

# 上传优化
```
sed -i "s|.d-flex {|.d-flex {font-size: 1px;|g" /data/seafile-server-7.0.5/seahub/media/css/seafile-ui.css
sed -i 's|("th",{width:"15%"},(0,f.gettext)("st|("th",{width:"25%"},(0,f.gettext)("st|g' /data/seafile-server-7.0.5/seahub/media/assets/frontend/js/app.faa7ba429394.js
```

* b34c01f55f48.css与app.faa7ba429394.js是混淆生成的,根据情况变
