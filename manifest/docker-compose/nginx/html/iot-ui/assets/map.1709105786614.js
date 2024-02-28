import{d as e,h as l,a,Y as t,S as n,n as o,ah as u,o as d,b as i,a8 as s,e as c,W as r,V as p,aA as v,aB as m}from"./vue.1709105786614.js";import{b as w,h as g,_ as f,H as V}from"./index.1709105786614.js";let y;const x={"class":"map-container"},b={"class":"coordinate-search"},h=(e=>(v("data-v-c2c221bc"),e=e(),m(),e))((()=>i("div",null,"-",-1))),k={key:0,"class":"address-result"};var _=f(e({__name:"map",emits:["updateMap"],setup(e,{expose:v,emit:m}){const f=l(null),_=l(""),C=l(""),M=l(""),B=l(""),F=l(""),L=l(!1),P=l(null);let G=null,U=null;const A=()=>{L.value=!1,I("updateMap",{lng:C.value,lat:M.value,address:_.value})},S=(e,l)=>{P.value&&(null==U||U.removeOverlay(P.value));const a=new G.Point(e,l);P.value=new G.Marker(a),null==U||U.addOverlay(new G.Marker(a)),null==U||U.setCenter(a),null==U||U.centerAndZoom(a,10)},Z=(e,l)=>{null==U||U.centerAndZoom(new G.Point(e,l),18);new G.Geocoder({extensions_town:!0}).getLocation(new G.Point(e,l),(function(e){e&&(_.value=e.content.poi_desc,B.value&&(_.value=B.value))}))},E=()=>{C.value&&M.value?(S(C.value,M.value),Z(C.value,M.value)):F.value&&(C.value="",M.value="",j(F.value))},j=e=>{if(e){const l=new G.LocalSearch(U);l.setSearchCompleteCallback((e=>{if(e){const l=e.getPoi(0);l&&(C.value=l.point.lng.toFixed(5),M.value=l.point.lat.toFixed(5),S(l.point.lng.toFixed(5),l.point.lat.toFixed(5)),Z(l.point.lng.toFixed(5),l.point.lat.toFixed(5)))}})),_.value=e,l.search(e||F.value)}},I=m;return v({openDialog:e=>{B.value="",L.value=!0,o((async()=>{const{BMapGL:l,centerPoint:a}=await new Promise(((e,l)=>{if(window.BMapGL)return e({BMapGL:window.BMapGL,centerPoint:y});Promise.all([w.getInfoByKey("sys.map.access.key"),w.getInfoByKey("sys.map.lngAndLat")]).then((([l,a])=>{const t=l.data.configValue,n=a.data.configValue;window.onBMapCallback=()=>{if(n){const[e,l]=n.split(",");y=new window.BMapGL.Point(e.trim(),l.trim())}else y="北京";e({BMapGL:window.BMapGL,centerPoint:y})};const o=document.createElement("script");o.type="text/javascript",o.src=`//api.map.baidu.com/api?v=1.0&type=webgl&ak=${t}&callback=onBMapCallback`,document.head.appendChild(o)}))["catch"]((()=>{l(new Error("地图加载失败，请刷新重试或联系开发者")),g.alert("地图加载失败，请刷新重试或联系开发者","提示",{type:"error"})}))}));G=l,U=new G.Map(f.value),e.lng&&e.lat?(C.value=e.lng,M.value=e.lat,E()):U.centerAndZoom(a,10),U.enableScrollWheelZoom(!0);const t=new G.ScaleControl;U.addControl(t);const n=new G.ZoomControl;U.addControl(n),U.addEventListener("click",(e=>(C.value=e.latlng.lng.toFixed(5),M.value=e.latlng.lat.toFixed(5),S(e.latlng.lng.toFixed(5),e.latlng.lat.toFixed(5)),Z(e.latlng.lng.toFixed(5),e.latlng.lat.toFixed(5)),!1))),U.addEventListener("zoomend",(()=>{const e=new G.Point(C.value,M.value),l=new G.Marker(e).getPosition();null==U||U.setCenter(l)})),e.address&&(B.value=e.address)})),C.value="",M.value="",F.value=""}}),(e,l)=>{const o=u("el-button"),v=u("el-input"),m=u("el-tooltip"),w=u("el-form-item"),g=u("el-dialog");return d(),a("div",null,[t(g,{title:"地图选点",modelValue:L.value,"onUpdate:modelValue":l[8]||(l[8]=e=>L.value=e),width:"900px","append-to-body":""},{"default":n((()=>[i("div",x,[i("div",b,[t(m,{"class":"box-item",effect:"dark",content:"点击放大镜或回车按键检索地址",placement:"top-start"},{"default":n((()=>[t(v,{modelValue:F.value,"onUpdate:modelValue":l[1]||(l[1]=e=>F.value=e),placeholder:"搜索地名",onKeyup:l[2]||(l[2]=s((e=>j(F.value)),["enter","native"]))},{append:n((()=>[t(o,{icon:c(V),onClick:l[0]||(l[0]=e=>j(F.value))},null,8,["icon"])])),_:1},8,["modelValue"])])),_:1}),t(v,{modelValue:C.value,"onUpdate:modelValue":l[3]||(l[3]=e=>C.value=e),placeholder:"经度"},null,8,["modelValue"]),h,t(v,{modelValue:M.value,"onUpdate:modelValue":l[4]||(l[4]=e=>M.value=e),placeholder:"纬度"},null,8,["modelValue"]),t(o,{onClick:E,type:"primary"},{"default":n((()=>[r("搜索")])),_:1})]),i("div",{"class":"map",ref_key:"mapContainer",ref:f},null,512),_.value?(d(),a("div",k,[t(w,{label:"经度","class":"input-item"},{"default":n((()=>[t(v,{modelValue:C.value,"onUpdate:modelValue":l[5]||(l[5]=e=>C.value=e)},null,8,["modelValue"])])),_:1}),t(w,{label:"纬度","class":"input-item"},{"default":n((()=>[t(v,{modelValue:M.value,"onUpdate:modelValue":l[6]||(l[6]=e=>M.value=e)},null,8,["modelValue"])])),_:1}),t(w,{label:"详细地址","class":"input-item"},{"default":n((()=>[t(v,{modelValue:_.value,"onUpdate:modelValue":l[7]||(l[7]=e=>_.value=e)},null,8,["modelValue"])])),_:1}),t(o,{onClick:A,style:{"margin-left":"10px"},type:"success"},{"default":n((()=>[r("确认")])),_:1})])):p("",!0)])])),_:1},8,["modelValue"])])}}}),[["__scopeId","data-v-c2c221bc"]]);export{_ as default};