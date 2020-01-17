// 当框框加载完成之后调用设置省份
window.onload = setSubject;
 
// 获取省市县/区的select选择框对象
var subject = document.getElementsByTagName("select")[0];
var year = document.getElementsByTagName("select")[1];
var term = document.getElementsByTagName("select")[2];
var unit = document.getElementsByTagName("select")[3];
var section = document.getElementsByTagName("select")[4];
 
// 设置省份
function setSubject() {
    // 遍历省份数组, provinceArr在city.js中
    for (var i = 0; i < subjectsArr.length; i++){
        // 创建省份option选项
        var opt = document.createElement("option");
        opt.value = subjectsArr[i];         // 设置value-提交给服务器用
        opt.innerHTML = subjectsArr[i];     // 设置option文本显示内容
        subject.appendChild(opt);          // 追加城市到下拉框
 
        // 当省份发生变化更改城市
    }
    subject.onchange = function(){
        setYear(this.selectedIndex);
    };
 
    // 省份加载完成，默认显示第一个省份的城市
    setYear(0);
}
 
// 设置城市
function setYear(subjectPos) {
    // 获取省份对象的城市，cityArr在city.js中
    var years  = yearsArr[subjectPos];
    year.length = 0;                  // 清空长度，重新从0开始赋值下拉框
 
    for (var i = 0; i < years.length; i++){
       // 创建城市option选项
       var opt = document.createElement("option");
       opt.value = years[i];         // 设置value-提交给服务器用
       opt.innerHTML = years[i];     // 设置option文本显示内容
        year.appendChild(opt);
    }
    year.onchange = function() {
        setTerm(subjectPos, this.selectedIndex);
    }
    // 默认显示城市的第一个县/区
    setTerm(subjectPos, 0);
}
 
// 设置县/区, 县/区是三位数组，需要传入哪个省份和城市
function setTerm(subjectPos, YearPos) {
    // 获取县/区，countyArr在city.js中国
    var terms = termsArr[subjectPos][YearPos];
    term.length = 0;
    
    for (var i = 0; i < terms.length; i++){
        // 创建县/区option选项
        var opt = document.createElement("option");
        opt.value = terms[i];         // 设置value-提交给服务器用
        opt.innerHTML = terms[i];     // 设置option文本显示内容
        term.appendChild(opt);        // 追加到县/区选择框中
    }
    term.onchange = function() {
        setUnit(subjectPos, YearPos,this.selectedIndex);
    }
    // 默认显示城市的第一个县/区
    setUnit(subjectPos, YearPos, 0);
}

function setUnit(subjectPos, YearPos, termPos) {
    // 获取县/区，countyArr在city.js中国
    var units = unitsArr[subjectPos][YearPos][termPos];
    unit.length = 0;
    
    for (var i = 0; i < units.length; i++){
        // 创建县/区option选项
        var opt = document.createElement("option");
        opt.value = units[i];         // 设置value-提交给服务器用
        opt.innerHTML = units[i];     // 设置option文本显示内容
        unit.appendChild(opt);        // 追加到县/区选择框中
    }
    unit.onchange = function() {
        setSection(subjectPos, YearPos,termPos,this.selectedIndex);
    }
    // 默认显示城市的第一个县/区
    setSection(subjectPos, YearPos,termPos,0);
}

function setSection(subjectPos, YearPos, termPos,unitPos) {
    // 获取县/区，countyArr在city.js中国
    var sections = sectionsArr[subjectPos][YearPos][termPos][unitPos];
    section.length = 0;
    
    for (var i = 0; i < sections.length; i++){
        // 创建县/区option选项
        var opt = document.createElement("option");
        opt.value = sections[i];         // 设置value-提交给服务器用
        opt.innerHTML = sections[i];     // 设置option文本显示内容
        section.appendChild(opt);        // 追加到县/区选择框中
    }
}