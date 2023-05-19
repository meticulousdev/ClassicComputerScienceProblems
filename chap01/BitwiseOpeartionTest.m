gene = 'AGCT';
for i = 1 : 9
    gene = append(gene, 'AGCT');
end

fprintf("gene          : %s\n", gene)
fprintf("gene length   : %d\n", length(gene))

bit_string = 1;

for i = 1 : length(gene)
    nucleotide = gene(i);
    bit_string = bitshift(bit_string, 2);

    if nucleotide == 'A'
        bit_string = bitor(bit_string, 0b00);
    elseif nucleotide == 'C'
        bit_string = bitor(bit_string, 0b01);
    elseif nucleotide == 'G'
        bit_string = bitor(bit_string, 0b10);
    elseif nucleotide == 'T'
        bit_string = bitor(bit_string, 0b11);
    else
        error("Invalid nucleotide: %s", nucleotide);
    end

    fprintf("%s", nucleotide)
    fprintf("%s", bit_string)
end