/**登录态*/
export const LOGINSTATE = false;
/**注册态*/
export const REGISTERSTATE = true;

/**表单类型*/
export interface FORM_TYPE {
    phoneNumber: string,
    account: string,
    password: string,
}
/**
 * 电话前缀
 */
export const countryCodes = [
    86, 60, 1, 81, 91, 886, 852, 355, 213, 93, 54, 353, 20, 251, 372, 971, 297,
    968, 376, 244, 1264, 1268, 61, 43, 994, 1246, 675, 1242, 375, 1441, 92, 595,
    973, 507, 359, 55, 229, 32, 354, 267, 48, 591, 501, 387, 975, 226, 257, 850,
    240, 45, 49, 670, 228, 1767, 1809, 593, 291, 7, 33, 298, 689, 679, 63, 358,
    238, 220, 243, 242, 299, 1473, 57, 506, 1671, 53, 592, 509, 82, 382, 31, 599,
    504, 233, 855, 241, 253, 420, 996, 263, 224, 245, 1345, 237, 974, 385, 269,
    254, 225, 965, 682, 266, 856, 371, 961, 231, 218, 423, 370, 40, 352, 250, 261,
    960, 356, 265, 223, 222, 389, 976, 880, 95, 51, 373, 212, 377, 258, 52, 264,
    27, 977, 505, 227, 234, 47, 351, 995, 46, 41, 503, 381, 232, 221, 357, 248,
    685, 966, 239, 1869, 1758, 378, 508, 1784, 94, 421, 386, 268, 249, 597, 252,
    66, 992, 676, 255, 1649, 1868, 90, 993, 216, 678, 502, 58, 673, 256, 380, 598,
    998, 34, 30, 65, 687, 64, 36, 963, 1876, 374, 967, 39, 964, 98, 62, 44, 1284,
    972, 962, 84, 260, 235, 350, 56, 236, 853,
];