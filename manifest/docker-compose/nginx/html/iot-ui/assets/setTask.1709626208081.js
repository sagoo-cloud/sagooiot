import{a as l}from"./index.170962620808141.js";import{_ as a,E as e}from"./index.1709626208081.js";import{d as s,a1 as t,h as o,k as c,aa as i,a as n,Y as d,S as r,ah as u,o as m,b as v,W as p,X as g,V as f,R as b,aA as h,aB as _}from"./vue.1709626208081.js";const w=s({name:"systemAddUser",setup(){const a=t({isShowDialog:!1,ruleForm:{target:"",uri:"",state:1,object:"",get_time:""},status:0,item_code:"",isShow:!1,testRes:""}),s=o("default"),n=o(),d=t({uri:[{required:!0,message:"请输入URL",trigger:"blur"}],object:[{required:!0,message:"请输入取值项",trigger:"blur"}],get_time:[{required:!0,message:"请输入取值周期",trigger:"blur"}]}),r=s=>{l.addDataSourceInfo(s).then((()=>{e.success("数据提交成功"),v(),a.isShow=!1}))},u=s=>{l.editataSourceInfo(s).then((l=>{e.success("数据提交成功"),v(),a.isShow=!1}))},m=(e,s)=>{let t={item_code:s,target_name:e.name};a.item_code=s,l.getDataSourceInfo(t).then((l=>{l?(a.ruleForm=l,a.status=1):(a.ruleForm={target:e.name,uri:"",state:1,object:"",get_time:""},a.status=2)}))},v=()=>{a.isShowDialog=!1};return c((()=>{})),{rules:d,openDialog:(l,e)=>{a.isShowDialog=!0,a.testRes="",m(l,e)},closeDialog:v,onCancel:()=>{v()},onSubmit:async l=>{l&&await l.validate((l=>{if(l)if(1===a.status)u(a.ruleForm);else{let l={...a.ruleForm,item_code:a.item_code};r(l)}}))},getDataSourceInfo:m,addDataSourceInfo:r,editataSourceInfo:u,test:()=>{let e={uri:a.ruleForm.uri,object:a.ruleForm.object};l.testDataSource(e).then((l=>{a.testRes=l}))},...i(a),formSize:s,ruleFormRef:n}}}),S=l=>(h("data-v-3a53246c"),l=l(),_(),l),F={"class":"system-add-user-container"},y={key:0,"class":"ico_down"},x={key:1,"class":"ico_up"},V={key:0,"class":"help-wrap"},R=[S((()=>v("div",{"class":"help-item"},[
v("div",{"class":"help-item-label"},"CRON表达式"),
v("div",{"class":"help-item-content"},"取值周期填写说明 eg:0 0 0 1 * * 每月一号执行一次")],-1))),S((()=>v("div",{"class":"help-item"},[
v("div",{"class":"help-item-label"},"CRON字段"),
v("div",{"class":"help-item-content"},[
v("div",{"class":"ant-row"},[
v("div",{"class":"ant-col ant-col-6"},"字段"),
v("div",{"class":"ant-col ant-col-6"},"允许值"),
v("div",{"class":"ant-col ant-col-6"},"允许特殊字符")]),
v("div",{"class":"ant-row"},[
v("div",{"class":"ant-col ant-col-6"},"-------------------"),
v("div",{"class":"ant-col ant-col-6"},"-------------------"),
v("div",{"class":"ant-col ant-col-6"},"-------------------")]),
v("div",{"class":"ant-row"},[
v("div",{"class":"ant-col ant-col-6"},"Seconds"),
v("div",{"class":"ant-col ant-col-6"},"0-59"),
v("div",{"class":"ant-col ant-col-6"},"* / , -")]),
v("div",{"class":"ant-row"},[
v("div",{"class":"ant-col ant-col-6"},"Minutes"),
v("div",{"class":"ant-col ant-col-6"},"0-59"),
v("div",{"class":"ant-col ant-col-6"},"* / , -")]),
v("div",{"class":"ant-row"},[
v("div",{"class":"ant-col ant-col-6"},"Hours"),
v("div",{"class":"ant-col ant-col-6"},"0-23"),
v("div",{"class":"ant-col ant-col-6"},"* / , -")]),
v("div",{"class":"ant-row"},[
v("div",{"class":"ant-col ant-col-6"},"Day"),
v("div",{"class":"ant-col ant-col-6"},"1-31"),
v("div",{"class":"ant-col ant-col-6"},"* / , - ?")]),
v("div",{"class":"ant-row"},[
v("div",{"class":"ant-col ant-col-6"},"Month"),
v("div",{"class":"ant-col ant-col-6"},"1-12 or JAN-DEC"),
v("div",{"class":"ant-col ant-col-6"},"* / , -")]),
v("div",{"class":"ant-row"},[
v("div",{"class":"ant-col ant-col-6"},"Week"),
v("div",{"class":"ant-col ant-col-6"},"0-6 or SUN-SAT"),
v("div",{"class":"ant-col ant-col-6"},"* / , - ?")])])],-1))),S((()=>v("div",{"class":"help-item"},[
v("div",{"class":"help-item-label"},"特殊字符"),
v("div",{"class":"help-item-content"},[
p(' "*"：当前字段所有的值   eg：第二字段*表示每分钟'),
v("br"),
p('"/"：描述范围的增量    eg：第二字段0-59/15表示每15分钟'),
v("br"),
p('","：分隔列表的项目    eg：第五字段1,3表示每周1,3执行'),
v("br"),
p('"-":连字符用于定义范围   eg：第三字段1-3表每天1到3点（含3点）'),
v("br"),
p('"?":设置当前字段为空'),
v("br")])],-1))),S((()=>v("div",{"class":"help-item"},[
v("div",{"class":"help-item-label"},"预定义"),
v("div",{"class":"help-item-content"},[
v("div",{"class":"ant-row"},[
v("div",{"class":"ant-col ant-col-8"},"字段"),
v("div",{"class":"ant-col ant-col-16"},"说明")]),
v("div",{"class":"ant-row"},[
v("div",{"class":"ant-col ant-col-8"},"----------------------------"),
v("div",{"class":"ant-col ant-col-16"},"---------------------------")]),
v("div",{"class":"ant-row"},[
v("div",{"class":"ant-col ant-col-8"},"@yearly or @annually"),
v("div",{"class":"ant-col ant-col-16"},"每年1月1日午夜运行一次")]),
v("div",{"class":"ant-row"},[
v("div",{"class":"ant-col ant-col-8"},"@monthly"),
v("div",{"class":"ant-col ant-col-16"},"每月一次，每月午夜运行一次")]),
v("div",{"class":"ant-row"},[
v("div",{"class":"ant-col ant-col-8"},"@weekly"),
v("div",{"class":"ant-col ant-col-16"},"每周运行一次，在星期六/星期日之间的午夜")]),
v("div",{"class":"ant-row"},[
v("div",{"class":"ant-col ant-col-8"},"@daily or @midnight"),
v("div",{"class":"ant-col ant-col-16"},"每天半夜运行")]),
v("div",{"class":"ant-row"},[
v("div",{"class":"ant-col ant-col-8"},"@hourly"),
v("div",{"class":"ant-col ant-col-16"},"每小时运行一次")])])],-1)))],k={"class":"dialog-footer"};var D=a(w,[["render",function(l,a,e,s,t,o){const c=u("el-form-item"),i=u("el-col"),h=u("el-input"),_=u("el-radio"),w=u("el-radio-group"),S=u("el-button"),D=u("el-row"),j=u("el-form"),C=u("el-dialog");return m(),n("div",F,[d(C,{title:"数据源配置接口",modelValue:l.isShowDialog,"onUpdate:modelValue":a[6]||(a[6]=a=>l.isShowDialog=a),width:"650px"},{footer:r((()=>[v("span",k,[d(S,{onClick:l.onCancel},{"default":r((()=>[p("取 消")])),_:1},8,["onClick"]),d(S,{type:"primary",onClick:a[5]||(a[5]=a=>l.onSubmit(l.ruleFormRef))},{"default":r((()=>[p("保 存")])),_:1})])])),"default":r((()=>[d(j,{model:l.ruleForm,ref:"ruleFormRef",rules:l.rules,"label-width":"90px"},{"default":r((()=>[d(D,{gutter:35},{"default":r((()=>[d(i,{xs:24,sm:24,md:24,lg:24,xl:24,"class":"mb20"},{"default":r((()=>[d(c,{label:"指标名称:"},{"default":r((()=>[p(g(l.ruleForm.target),1)])),_:1})])),_:1}),d(i,{xs:24,sm:24,md:24,lg:24,xl:24,"class":"mb20"},{"default":r((()=>[d(c,{label:"URL:",prop:"uri"},{"default":r((()=>[d(h,{modelValue:l.ruleForm.uri,"onUpdate:modelValue":a[0]||(a[0]=a=>l.ruleForm.uri=a),placeholder:"请输入URL",clearable:""},null,8,["modelValue"])])),_:1})])),_:1}),d(i,{xs:24,sm:24,md:24,lg:24,xl:24,"class":"mb20"},{"default":r((()=>[d(c,{label:"取值项:",prop:"object"},{"default":r((()=>[d(h,{modelValue:l.ruleForm.object,"onUpdate:modelValue":a[1]||(a[1]=a=>l.ruleForm.object=a),placeholder:"请输入取值项",clearable:""},null,8,["modelValue"])])),_:1})])),_:1}),d(i,{xs:24,sm:24,md:24,lg:24,xl:24,"class":"mb20"},{"default":r((()=>[d(c,{"class":"inline-row",label:"取值周期:",prop:"get_time"},{"default":r((()=>[d(h,{modelValue:l.ruleForm.get_time,"onUpdate:modelValue":a[2]||(a[2]=a=>l.ruleForm.get_time=a),placeholder:"请输入取值周期",clearable:""},null,8,["modelValue"]),v("div",{"class":"tip",onClick:a[3]||(a[3]=a=>l.isShow=!l.isShow)},[l.isShow?(m(),n("span",x)):(m(),n("span",y)),p(" 帮助 ")])])),_:1})])),_:1}),l.isShow?(m(),n("div",V,R)):f("",!0),d(i,{xs:24,sm:24,md:24,lg:24,xl:24,"class":"mb20"},{"default":r((()=>[d(c,{label:"是否启用:"},{"default":r((()=>[d(w,{modelValue:l.ruleForm.state,"onUpdate:modelValue":a[4]||(a[4]=a=>l.ruleForm.state=a),"class":"ml-4"},{"default":r((()=>[d(_,{size:"large",label:1},{"default":r((()=>[p("启用")])),_:1}),d(_,{size:"large",label:2},{"default":r((()=>[p("禁用")])),_:1})])),_:1},8,["modelValue"]),l.ruleForm.uri&&l.ruleForm.object?(m(),b(S,{key:0,onClick:l.test,style:{"margin-left":"20px"},size:"small",type:"primary"},{"default":r((()=>[p("检测")])),_:1},8,["onClick"])):f("",!0)])),_:1})])),_:1}),l.testRes||l.testRes.toString()?(m(),b(i,{key:1,xs:24,sm:24,md:24,lg:24,xl:24,"class":"mb20"},{"default":r((()=>[d(c,{label:"测试结果:"},{"default":r((()=>[v("span",null,"数据源返回数据值:"+g(l.testRes),1)])),_:1})])),_:1})):f("",!0)])),_:1})])),_:1},8,["model","rules"])])),_:1},8,["modelValue"])])}],["__scopeId","data-v-3a53246c"]]);export{D as default};