classdef CompressedGene
    properties
        bit_string
    end

    methods
        function obj = CompressedGene(gene)
            obj.compress(gene);
        end

        function obj = compress(obj, gene)
            obj.bit_string = 1;

            for i = 1 : length(gene)
                nucleotide = gene(i);
                obj.bit_string = bitshift(obj.bit_string, 2);
                
                if nucleotide == 'A'
                    obj.bit_string = bitor(obj.bit_string, 0b00);
                elseif nucleotide == 'C'
                    obj.bit_string = bitor(obj.bit_string, 0b01);
                elseif nucleotide == 'G'
                    obj.bit_string = bitor(obj.bit_string, 0b10);
                elseif nucleotide == 'T'
                    obj.bit_string = bitor(obj.bit_string, 0b11);
                else
                    error("Invalid nucleotide: %s", nucleotide);
                end
            end
        end

        function gene = decompress(obj)
            gene = '';
            for i = 0 : 2 : length(dec2bin(obj.bit_string))
                bits = bitand((bitshift(obj.bit_string, i)), 0b11);
                
                if bits == 0b00
                    gene = append(gene, 'A');
                elseif bits == 0b01
                    gene = append(gene, 'C');
                elseif bits == 0b10
                    gene = append(gene, 'G');
                elseif bits == 0b11
                    gene = append(gene, 'T');
                end
            end
        end
    end
end