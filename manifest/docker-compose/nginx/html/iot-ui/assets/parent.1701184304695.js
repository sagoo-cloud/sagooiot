import{L as h,a2 as v,l as i,K as d,n as w,H as y,i as g,ab as _,a as k,_ as m,T as f,j as V,aD as N,ai as R,o as r,$ as C,S as c,W as A,P as K,aG as T}from"./vue.1701184304695.js";import{_ as L,u as B}from"./index.1701184304695.js";const P=h({name:"layoutParentView",props:{minHeight:{type:String,default:""}},setup(){const{proxy:t}=V(),a=N(),o=B(),e=v({refreshRouterViewKey:null,keepAliveNameList:[]}),l=i(()=>o.state.themeConfig.themeConfig.animation),u=i(()=>o.state.themeConfig.themeConfig),s=i(()=>o.state.keepAliveNames.keepAliveNames);return d(()=>{e.keepAliveNameList=s.value,t.mittBus.on("onTagsViewRefreshRouterView",n=>{e.keepAliveNameList=s.value.filter(p=>a.name!==p),e.refreshRouterViewKey=null,w(()=>{e.refreshRouterViewKey=n,e.keepAliveNameList=s.value})})}),y(()=>{t.mittBus.off("onTagsViewRefreshRouterView")}),g(()=>a.fullPath,()=>{e.refreshRouterViewKey=a.fullPath}),{getThemeConfig:u,getKeepAliveNames:s,setTransitionName:l,..._(e)}}}),$={class:"height:100%",style:{"overflow-y":"auto","overflow-x":"hidden"}};function H(t,a,o,e,l,u){const s=R("router-view");return r(),k("div",$,[m(s,null,{default:f(({Component:n})=>[m(C,{name:t.setTransitionName,mode:"out-in"},{default:f(()=>[(r(),c(T,{include:t.keepAliveNameList},[(r(),c(A(n),{key:t.refreshRouterViewKey,class:"w100",style:K({minHeight:t.minHeight})},null,8,["style"]))],1032,["include"]))]),_:2},1032,["name"])]),_:1})])}var D=L(P,[["render",H]]);export{D as default};