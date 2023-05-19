gene = 'AGCT';
for i = 1 : 9
    gene = append(gene, 'AGCT');
end

fprintf("gene          : %s\n", gene);
fprintf("gene length   : %d\n", length(gene));

bit_string = 1;
fprintf("%d %s\n", bit_string, dec2bin(bit_string));

bit_string = bitshift(bit_string, 1);
fprintf("%d %s\n", bit_string, dec2bin(bit_string));

bit_string = 1;
bit_string_ret = bitor(bit_string, 0b00);
fprintf("%s | %s = %s\n", dec2bin(bit_string), dec2bin(0b00), dec2bin(bit_string_ret));
% 10 | 00 = 10

bit_string_ret = bitor(bit_string, 0b01);
fprintf("%s | %s = %s\n", dec2bin(bit_string), dec2bin(0b01), dec2bin(bit_string_ret));
% 10 | 01 = 11

bit_string_ret = bitor(bit_string, 0b10);
fprintf("%s | %s = %s\n", dec2bin(bit_string), dec2bin(0b10), dec2bin(bit_string_ret));
% 10 | 10 = 10

bit_string_ret = bitor(bit_string, 0b11);
fprintf("%s | %s = %s\n", dec2bin(bit_string), dec2bin(0b11), dec2bin(bit_string_ret));
% 10 | 11 = 11
