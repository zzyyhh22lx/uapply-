// 工具函数

const map: { [key: string]: string } = {
    '<' : '&lt;',
    '>' : '&gt;',
    '/': '&#x2F;',
    '&': '&amp;',
    "'": '&#x27;',
    '"': '&quot;'
};
/**
 * 转义，将特殊字符转义成unicode编码
 * @param str 
 * @returns 
 */
export function transformStr(str: string | undefined) {
    if(!str || typeof str !== 'string') return '';
    return str.replace(/[<>'/&"]/g, function(char) {
        return map[char];
    })
}


/**
 * / 超出字符隐藏(默认15)
 * @param value 内容
 * @param len 显示的最大字符
 * @returns 
 */
export function ellipsis(value = '', len = 15) {
    if (!value) return '';
    if (value.length > len) {
      return `${value.slice(0, len)}...`;
    };
    return value;
};
